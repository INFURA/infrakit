moved {
  from = aws_route53_zone.din_dev[0]
  to = module.route53_zone_din_dev[0].aws_route53_zone.zone
}
