kubectl run flink-sql-client --rm -it --env=FLINK_PROPERTIES="jobmanager.rpc.address: flink-jobmanager.tools" --image=apache/flink:1.15.0-scala_2.12 -n tools --restart=Never -- bin/sql-client.sh