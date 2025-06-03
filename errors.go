package hibiscus

import "errors"

// 身份认证 & 注册相关错误

var ErrEmptyPassword = errors.New("empty password") // 密码为空，通常用于注册或登录校验失败

var ErrEmptyEmail = errors.New("empty email") // 邮箱为空，通常用于注册、登录、找回密码等操作

var ErrSameEmail = errors.New("same email") // 新旧邮箱相同，用户尝试更换邮箱时触发

var ErrEmailExists = errors.New("email exists, please use another email") // 邮箱已存在，尝试注册或更新为已被注册的邮箱

var ErrUserNotExists = errors.New("user not exists") // 用户不存在，常用于登录、查询或操作不存在的用户时

var ErrUnauthorized = errors.New("unauthorized") // 未授权访问，例如缺少登录状态或 token

var ErrForbidden = errors.New("forbidden access") // 拒绝访问，用户虽已登录但无权限访问目标资源

var ErrUserNotAllowLogin = errors.New("user not allow login") // 用户被禁止登录，可能是被管理员封禁

var ErrUserNotAllowSignup = errors.New("user not allow signup") // 用户被禁止注册，系统配置或策略限制注册行为

var ErrNotActivated = errors.New("user not activated") // 用户账户未激活，通常用于邮箱激活未完成

var ErrTokenRequired = errors.New("token required") // 缺少必要的令牌，例如访问受保护资源时

var ErrInvalidToken = errors.New("invalid token") // 令牌格式非法或不符合规范

var ErrBadToken = errors.New("bad token") // 令牌已被篡改、伪造或无效

var ErrTokenExpired = errors.New("token expired") // 令牌已过期

var ErrEmailRequired = errors.New("email required") // 邮箱字段必须提供但未提供

// 通用资源/数据处理相关错误

var ErrNotFound = errors.New("not found") // 请求的数据或资源未找到

var ErrNotChanged = errors.New("not changed") // 数据未发生变化，例如更新请求中没有实际变更字段

var ErrInvalidView = errors.New("with invalid view") // 请求使用了无效的视图标识或参数

// 权限与逻辑控制相关错误

var ErrOnlySuperUser = errors.New("only super user can do this") // 仅限超级用户执行的操作

var ErrInvalidPrimaryKey = errors.New("invalid primary key") // 主键非法，可能为格式错误或缺失
