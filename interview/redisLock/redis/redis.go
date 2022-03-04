package redis

// redis 客户端以及一些操作
// 参考：https://github.com/bsm/redislock

/*
## 前言
回望过去数年，大概待过 4 个公司，基本上每家都用 redis，当然 MySQL 也使用，这足见 redis 在各个公司的技术栈中已经是不可或缺的一个。
尽管如此，对 redis 的了解也很局限，大都停留在“缓存”上，而实际上， redis 可以在很多场景被使用。从今天开始，我们尝试 redis 的一些不一样的用法。
先从 redis 的分布式锁开始。

## 分布式锁

*/
