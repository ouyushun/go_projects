need_start_server_shell=(
  # rpc
  restart-user-rpc-test.sh

  # api
)

for i in ${need_start_server_shell[*]} ; do
    chmod +x $i
    ./$i
done


docker ps