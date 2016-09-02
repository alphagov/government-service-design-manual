# encoding: utf-8
require 'sinatra/base'
require 'gds_api/rummager'
require 'liquid'

LANG="en_US.UTF-8"
LC_ALL="en_US.UTF-8"

class ServiceManual < Sinatra::Base
  Tilt.register Tilt::LiquidTemplate, :html

  before do
    cache_control :public, :must_revalidate, :max_age => 3600
  end

  get '/' do
    redirect '/service-manual'
  end

  get '/service-manual/search' do
    Liquid::Template.file_system = Liquid::LocalFileSystem.new(File.join(settings.root, '_search', 'includes'))

    site_settings = { 'govuk_template_assets' => '/service-manual/assets/govuk_template' }
    page_settings = { 'header_class' => 'with-proposition', 'title' => 'Search' }

    @search_term = params[:q]

    begin
      if @search_term.nil? or @search_term.strip.empty?
        @results = []
      else
        res = search_client.search(q: @search_term, filter_manual: "service-manual")
        @results = res['results'].map { |result|
          result.merge({'title' => result['title'].gsub(/\AGovernment Service Design Manual: /, '')})
        }
        @count = res['total']
        page_settings['title'] = "Search results for #{@search_term}"
      end
    rescue GdsApi::BaseError => e
      status 500
      @results = []
    end

    liquid :search,
           :views => '_search',
           :layout => :govuk,
           :layout_options => { :views => '_search/layouts' },
           :locals => { :query => @search_term, :count => @count, :results => @results, :site => site_settings, :page => page_settings }
  end

  get '*.html' do
    request_path_without_suffix = params[:splat][0]
    redirect request_path_without_suffix
  end

  get '*/index' do
    request_path_without_trailing_index = params[:splat][0]
    redirect request_path_without_trailing_index
  end

  get '*' do
    jekyll_root = '_site'
    request_path = params[:splat][0]

    index_file = File.join(jekyll_root, request_path, 'index.html')
    html_file = File.join(jekyll_root, request_path + '.html')
    static_file = File.join(jekyll_root, request_path)

    if FileTest.file? index_file
      send_file index_file
    elsif FileTest.file? html_file
      send_file html_file
    elsif FileTest.file? static_file
      send_file static_file
    else
      status 404
      return "No file exists to serve #{request_path}"
    end
  end

  helpers do
    def asset_path(filename)
      if filename.end_with? '.css'
        "/service-manual/assets/govuk_template/stylesheets/#{filename}"
      elsif filename.end_with? '.js'
        "/service-manual/assets/govuk_template/javascripts/#{filename}"
      elsif filename.end_with? '.png' or filename.end_with? '.ico'
        "/service-manual/assets/govuk_template/images/#{filename}"
      else
        ""
      end
    end
  end

  def search_client
    @search_client ||= GdsApi::Rummager.new(Plek.current.find('rummager'))
  end

  # Start the server if this Ruby file is executed directly
  run! if app_file == $0
end
