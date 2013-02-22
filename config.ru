require "rack-slashenforce"

USERNAME = ENV['AUTH_USERNAME']
PASSWORD = ENV['AUTH_PASSWORD']

use Rack::Auth::Basic, "Restricted" do |username, password|
  [username, password] == [USERNAME, PASSWORD]
end

use Rack::AppendTrailingSlash

use Rack::Static, :urls => [""], :root => "_site", :index => "index.html"

run lambda { |env|
  # Should never get here
  [500, {}, []]
}
