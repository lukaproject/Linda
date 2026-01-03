# pprof

我们已经给agentcentral注入了pprof组件，访问http://{serviceIP:servicePort}/debug/pprof/ 即可访问pprof组件。
但是要启用这个组件，请在环境变量中添加上env=debug （当然现阶段你也可以不加，默认启动就是debug模式）