require 'rack/jekyll'

run Rack::Jekyll.new

USERNAME = ENV['AUTH_USERNAME']
PASSWORD = ENV['AUTH_PASSWORD']
use Rack::Auth::Basic, "Restricted" do |username, password|
  [username, password] == [USERNAME, PASSWORD]
end