# KeyZenDB



## Feature
- [ ] Concurrency: Lock-Free Data Structures (redis),Optimistic Concurrency Control (OCC)
چون In-Memory DB برای سرعت بالا ساخته شده، باید بتونه هم‌زمان چندین درخواست رو مدیریت کنه بدون اینکه داده‌ها خراب بشن.

- [ ] Atomicity: Atomic Commands
وقتی عملیاتی روی داده انجام می‌شه، یا باید کاملاً انجام بشه یا اصلاً انجام نشه.
مثلاً اگه داری یه مقدار رو به یه متغیر اضافه می‌کنی، نباید نصفه‌نیمه ذخیره بشه.

- [ ] Caching: Time-To-Live (TTL),Least Recently Used (LRU) Eviction
- [ ] Replication: Master-Slave Replication ,Leader-Follower Model
https://redis.io/docs/latest/operate/oss_and_stack/management/replication/
- [ ] Persistence: RDB,AOF,No persistence
https://redis.io/docs/latest/operate/oss_and_stack/management/persistence/
- [ ] Scalability: Load Balancing, Cluster Mode
- [ ] High availability : 
https://redis.io/docs/latest/operate/oss_and_stack/management/sentinel/