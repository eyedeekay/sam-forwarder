#
# Regular cron jobs for the sam-forwarder package
#
0 4	* * *	root	[ -x /usr/bin/sam-forwarder_maintenance ] && /usr/bin/sam-forwarder_maintenance
