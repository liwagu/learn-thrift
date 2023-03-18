namespace go example

struct UserProfile {
  1: i32 id,
  2: string name,
  3: string email,
}

service UserService {
  UserProfile getUserProfile(1: i32 userId),
}

// 1.1 namespace go example 用于生成go代码时，生成的代码在包名为example中，也就是生成的go代码为：
// package example
// 1.2 struct UserProfile {
//   1: i32 id,
//   2: string name,
//   3: string email,
// } 
// 定义了一个结构体UserProfile，其中id，name和email分别为int32，string，string类型
// 1.3 service UserService {
//   UserProfile getUserProfile(1: i32 userId),
// } 
// 定义了一个服务UserService，该服务内有一个方法getUserProfile，该方法的参数为int32类型的userId，返回值为UserProfile结构体类型。