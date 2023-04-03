# openiot

Four main modules: 
- user
  - collection ? 不知道是什么功能，应该是用户组之类的
- admin
- node 垃圾桶节点，之前被叫做 Bin
  - level
  > Info 不要简写成 Inf 啊，看不出来我还以为是 Interface
- blog 文章，原名 essay **不需要了，砍掉！**
  - tag
  - category
- comment 评论 **同文章，舍弃**
- communication (mail -> ?)
- utils
  - weather
  - address/location
  
# Creative
Primarily, this is a management system to help manager charge to the status of factory.
Therefore, the platform will be used by many people and need to divide the authority of users, to protect data from change of unkind propose.
There are two tables and a middle table for this feature, 
- user: the main user features including login, check info
- role: the roles with different authority
  - divide into three roles and root: 
    - manager read and write
    - leader: only-read & leader domain change
    - worker: only-read
    - super role(root): program administrator
  > the domain of leader need to discuss
- authority: store the id2id relation of user and role.

Furthermore, the core function of robot dog(RD in follows) scheduling need a well design. The vital necessarily trouble is that RB is dynamic in the map and how to show their position in time. The first approach is the long connection of http protocol like websocket and RB update their position per 5 minutes. There are some relative problem that:
- can RB be controlled to went to a given address? (Maybe not)
- which data will RB update to server?
  > Industrial indicators such as temperature and humidity or content of industrial sites
  - need a lot of design
  - Customizable parameter settings? A given limit in init and reflect when warming.
  > 接触到了一些关于 PaaS 和 SaaS 的内容，很有趣，和我未来规划方向相符
  > 这里我们需要划分清楚核心功能，将其做为 PaaS，然后进行集成 SaaS 的方式进行定制化服务的开发，那么我们这里实际上需要设计的是原本的核心功能与接口，以及定制化服务的框架或者脚手架。这样子就很清楚了，user 和 authority 的部分是 core 的范围，但是允许设计外接的部分，比如在 User 中加入一个 UserExternal 的拓展接口，使得具体部署到的服务可以进行注入。那么我们维护的就是两个分支，分别是核心功能和脚手架，以及具体应用 demo
  > https://www.jianshu.com/p/55f2cbdbead0
  > https://blog.csdn.net/zyzBulus/article/details/81169733
  > 参考这个链接的 external 设计，抽象一个表用于存储机械狗监控的指标数据。可以叫做 quota ？
  - Send a message to admins phones. use the service of SMS.
- all RB patrol in the factory and 