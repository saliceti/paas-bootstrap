#!/usr/bin/env ruby

require 'optparse'
require 'yaml'
require File.expand_path("../../../shared/lib/secret_generator", __FILE__)

generator = SecretGenerator.new(
  "bosh_postgres_password" => :simple,
  "bosh_mbus_bootstrap_password" => :simple,
  "bosh_registry_password" => :simple,
  "bosh_redis_password" => :simple,
  "bosh_hm_director_password" => :simple,
  "bosh_admin_password" => :simple,
  "bosh_bosh_exporter_password" => :simple,
  "bosh_vcap_password" => :sha512_crypted,
  "bosh_uaa_postgres_password" => :simple,
  "bosh_credhub_postgres_password" => :simple,
  "bosh_uaa_login_client_password" => :simple,
  "bosh_uaa_uaa_encryption_key_1" => :simple,
  "bosh_uaa_jwt_signing_key" => :bosh_rsa_key,
)

option_parser = OptionParser.new do |opts|
  opts.on('--existing-secrets FILE') do |file|
    existing_secrets = YAML.load_file(file)
    # An empty file parses as false
    generator.existing_secrets = existing_secrets["secrets"] if existing_secrets
  end
end
option_parser.parse!

output = { "secrets" => generator.generate }
puts output.to_yaml
