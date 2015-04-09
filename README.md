gearman


监听端口
worker注册
函数注册
请求分发、结果返回


启动 Server ,监听 4732
启动 Worker , 注册功能到服务器上.
启动 Client ,请求注册的功能,并提供输入的数据
Gearman 的服务发送"功能请求"到 Worker.并给 Client 输入的数据传送给 Worker.
Worker 接收到"功能请求"和输入的数据.并处理输入的数据内容.
Worker 返回处理的数据结果,服务转发这个结果给客户端.
Client 客户端接收到处理的结果,并显示结果给用户


A job server liked gearman and written by go 


server  
client  send job
worker  exec job


