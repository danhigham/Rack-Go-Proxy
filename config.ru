require './master.rb'

# kick off go slave
`chmod +x ./slave`

# start rack master
run RackMaster.new