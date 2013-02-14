require 'socket'

class RackMaster

  def initialize()
    @master, slave = Socket.pair(:UNIX, :DGRAM, 0)
    pid = Process.spawn('./slave', slave.fileno.to_s, { :close_others => false })
    slave.close
  end

  def call(env)
    req = Rack::Request.new(env)
    uri = env["REQUEST_URI"]

    @master.puts uri

    resp = @master.gets
    [200, {"Content-Type" => "text/plain"}, [resp]]
  end

end