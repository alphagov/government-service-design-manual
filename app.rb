require 'sinatra/base'

class ServiceManual < Sinatra::Base
  get '/' do
    redirect '/service-manual'
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

  # Start the server if this Ruby file is executed directly
  run! if app_file == $0
end
