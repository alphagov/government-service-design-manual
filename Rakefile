namespace :router do
  task :router_environment do
    require 'logger'
    require 'plek'
    require 'gds_api/router'
    @logger = Logger.new STDOUT
    @logger.level = Logger::DEBUG

    @router_api = GdsApi::Router.new(Plek.current.find('router-api'))
    @app_id = 'service-manual'
  end

  task :register_backend => :router_environment do
    url = Plek.current.find(@app_id, :force_http => true) + "/"
    @logger.info "Registering #{@app_id} application against #{url}"
    @router_api.add_backend(@app_id, url)
  end

  task :register_routes => [ :router_environment ] do
    [
      %w(/service-manual prefix),
    ].each do |path, type|
      begin
        @logger.info "Registering #{type} route #{path}"
        @router_api.add_route(path, type, @app_id)
      rescue => e
        @logger.error "Error registering route: #{e.message}"
        raise
      end
    end

    [
      # Missed redirections from service manual (post-beta only)
      ["/service-manual/browsers-and-devices", "exact", "/service-manual/user-centred-design/browsers-and-devices"],
      ["/service-manual/browsers-and-devices/index", "exact", "/service-manual/user-centred-design/browsers-and-devices"],
      ["/service-manual/browsers-and-devices/index.html", "exact", "/service-manual/user-centred-design/browsers-and-devices"],
      ["/service-manual/content-and-design/browsers-and-devices", "exact", "/service-manual/user-centred-design/browsers-and-devices"],
      ["/service-manual/content-and-design/browsers-and-devices.html", "exact", "/service-manual/user-centred-design/browsers-and-devices"],
      ["/service-manual/design-and-content/browsers-and-devices", "exact", "/service-manual/user-centred-design/browsers-and-devices"],
      ["/service-manual/design-and-content/browsers-and-devices.html", "exact", "/service-manual/user-centred-design/browsers-and-devices"],

      # Redirections for old content
      ["/service-manual/digital-by-default/assessments-before-2014", "exact", "/service-manual/digital-by-default"],
      ["/service-manual/digital-by-default/assessments-before-2014.html", "exact", "/service-manual/digital-by-default"],
      ["/service-manual/digital-by-default/go-live-panel", "exact", "/service-manual/digital-by-default"],
      ["/service-manual/digital-by-default/go-live-panel.html", "exact", "/service-manual/digital-by-default"],
      ["/service-manual/identity-assurance/guidance-for-government-service-providers", "exact", "/service-manual/identity-assurance"],
      ["/service-manual/identity-assurance/guidance-for-government-service-providers.html", "exact", "/service-manual/identity-assurance"],
      ["/service-manual/making-software/accessibility-testing", "exact", "/service-manual/user-centred-design/user-research/accessibility-testing"],
      ["/service-manual/making-software/accessibility-testing.html", "exact", "/service-manual/user-centred-design/user-research/accessibility-testing"],
      ["/service-manual/the-team/induction-and-development", "exact", "/service-manual/the-team/learning-and-development"],
      ["/service-manual/the-team/induction-and-development.html", "exact", "/service-manual/the-team/learning-and-development"],
      ["/service-manual/user-centered-design", "prefix", "/service-manual/user-centred-design"],
      ["/service-manual/user-centred-design/service-look-and-feel", "exact", "/service-manual/user-centred-design/service-user-experience"],
      ["/service-manual/user-centred-design/service-look-and-feel.html", "exact", "/service-manual/user-centred-design/service-user-experience"],
      ["/service-manual/user-centered-design/user-research/guerilla-testing", "prefix", "/service-manual/user-centred-design/user-research/guerrilla-testing"],
      ["/service-manual/user-centered-design/user-research/guerrilla-testing", "prefix", "/service-manual/user-centred-design/user-research/guerrilla-testing"],
      ["/service-manual/user-centered-design/what-should-service-look-like", "exact", "/service-manual/user-centred-design/service-user-experience"],
      ["/service-manual/user-centered-design/what-should-service-look-like.html", "exact", "/service-manual/user-centred-design/service-user-experience"],
      ["/service-manual/user-centered-design/writing-government-services", "exact", "/service-manual/content-designers/transactions-style-guide"],
      ["/service-manual/user-centered-design/writing-government-services.html", "exact", "/service-manual/content-designers/transactions-style-guide"],
      ["/service-manual/user-centred-design/know-your-users", "exact", "/service-manual/user-centred-design/user-needs"],
      ["/service-manual/user-centred-design/know-your-users.html", "exact", "/service-manual/user-centred-design/user-needs"],
      ["/service-manual/user-centred-design/introduction-to-user-research", "exact", "/service-manual/user-centred-design/user-research"],
      ["/service-manual/user-centred-design/introduction-to-user-research.html", "exact", "/service-manual/user-centred-design/user-research"],
    ].each do |path, type, destination|
      begin
        @logger.info "Registering #{type} redirect route from #{path} -> #{destination}"
        @router_api.add_redirect_route(path, type, destination, "permanent")
      rescue => e
        @logger.error "Error registering route: #{e.message}"
        raise
      end
    end
    @router_api.commit_routes
  end

  desc "Register the service manual backend and routes with the GOV.UK router"
  task :register => [ :register_backend, :register_routes ]
end

namespace :rummager do
  task :index do
    require 'plek'
    require 'rummageable'

    search_index = File.expand_path('../.search-index.json', __FILE__)
    puts "Registering search index: #{search_index}"

    content = File.read(search_index)
    parsed = JSON.parse(content)
    puts "Handling #{parsed.count} records from #{search_index}"

    rummageable_index = Rummageable::Index.new(Plek.current.find('rummager'), '/service-manual')
    rummageable_index.add_batch(parsed)
  end
end
