require "rack-rewrite"

USERNAME = ENV['AUTH_USERNAME']
PASSWORD = ENV['AUTH_PASSWORD']

use Rack::Auth::Basic, "Restricted" do |username, password|
  [username, password] == [USERNAME, PASSWORD]
end

use Rack::Rewrite do
  # Rewrite any requests that contain no dots and end without a trailing slash
  # to contain a trailing slash. This makes Rack::Static do the right thing
  # for directories, and enforces consistency with the style of the rest of
  # GOV.UK.
  rewrite %r{^([^.]+[^/])$}, '$1/'
  # Redirect any requests that end with a slash to the equivalent without the
  # trailing slash.
  r301 %r{^(.+?)/+$}, '$1'
end

use Rack::Static, :urls => [""], :root => "_site", :index => "index.html"

run lambda { |env|
  # Should never get here
  [500, {}, []]
}
