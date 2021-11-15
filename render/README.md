# README

## render模块中

- 通过调用RenderTask RenderTag来调用对应的Renderer
- 实现原理是调用对应的Formattor，在Formatter中根据用户的打印设置来决定记录的字符串内容
