package cmd

import (
	"os"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	infrakitv1alpha1 "github.com/INFURA/infrakit/api/v1alpha1"
	"github.com/INFURA/infrakit/controllers"
	"github.com/INFURA/infrakit/controllers/podtemplateclientset"
	"github.com/INFURA/infrakit/controllers/proxy"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	//+kubebuilder:scaffold:imports
)

type OperatorCmdConfig struct {
	MetricsBindAddress     string `koanf:"metrics-bind-address"`
	HealthProbeBindAddress string `koanf:"health-probe-bind-address"`
	LeaderElect            bool   `koanf:"leader-elect"`
}

var (
	scheme            = runtime.NewScheme()
	setupLog          = ctrl.Log.WithName("setup")
	operatorCmdConfig OperatorCmdConfig
)

func init() {
	rootCmd.AddCommand(operatorCmd)

	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(infrakitv1alpha1.AddToScheme(scheme))
	//+kubebuilder:scaffold:scheme
}

var operatorCmd = &cobra.Command{
	Use:   "operator",
	Short: "Run the InfraKit Kubernetes operator",
	Run:   operatorMain,
}

func operatorMain(cmd *cobra.Command, args []string) {
	if err := k.Unmarshal("operator", &operatorCmdConfig); err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:                 scheme,
		MetricsBindAddress:     viper.GetString("metrics-bind-address"),
		Port:                   9443,
		HealthProbeBindAddress: operatorCmdConfig.HealthProbeBindAddress,
		LeaderElection:         operatorCmdConfig.LeaderElect,
		LeaderElectionID:       "2ac48f74.infura.io",
		// LeaderElectionReleaseOnCancel defines if the leader should step down voluntarily
		// when the Manager ends. This requires the binary to immediately end when the
		// Manager is stopped, otherwise, this setting is unsafe. Setting this significantly
		// speeds up voluntary leader transitions as the new leader don't have to wait
		// LeaseDuration time first.
		//
		// In the default scaffold provided, the program ends immediately after
		// the manager stops, so would be fine to enable this option. However,
		// if you are doing or is intended to do any operation such as perform cleanups
		// after the manager stops then its usage might be unsafe.
		// LeaderElectionReleaseOnCancel: true,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	if err = (&proxy.MainReconciler{Reconciler: controllers.Reconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Proxy")
		os.Exit(1)
	}
	if err = (&podtemplateclientset.MainReconciler{Reconciler: controllers.Reconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "PodTemplateClientSet")
		os.Exit(1)
	}
	//+kubebuilder:scaffold:builder

	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up health check")
		os.Exit(1)
	}
	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up ready check")
		os.Exit(1)
	}

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}
