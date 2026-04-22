# OpFlow — AI Content Workflow Engine (Go + Next.js)

> 多源数据采集 → LLM 推理分析 → 结构化内容合成的 Agent-like 工作流引擎。支持 SSE 流式推送、Prompt 工程优化与 Docker 微服务部署。

---

## 🚀 功能特性

### 核心能力
- **多源数据采集**：自动聚合知乎等平台热点内容，统一清洗与格式转换
- **LLM 推理与合成**：基于 Prompt 工程的内容分析与结构化输出
- **多平台适配**：支持不同场景下的内容格式生成
- **实时流式推送**：通过 SSE 技术实现生成进度实时推送
- **响应式界面**：适配各种设备的现代化用户界面

### 工程亮点
- **Agent-like 工作流**：Collector → Analyzer → Synthesizer 三节点编排
- **微服务架构**：模块化服务设计，便于扩展和维护
- **容器化部署**：完整的 Docker 容器化方案，一键部署
- **实时通信**：基于 SSE 的服务器推送技术，TTFB < 200ms
- **智能调度**：任务调度器支持定时与触发双模式

---

## Architecture

OpFlow 采用 **Collector → Analyzer → Synthesizer** 三节点工作流设计，支持异步任务编排与实时流式输出：

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│  Collector  │ --> │  Analyzer   │ --> │ Synthesizer │
│  多源采集    │     │  LLM推理    │     │ 内容合成    │
└─────────────┘     └─────────────┘     └─────────────┘
      ↑                                      ↓
   API/爬虫                                SSE推送
```

| 节点 | 职责 | 对应服务 |
|------|------|---------|
| **Collector** | 多源数据采集与清洗 | `crawler` + `hotspot API` |
| **Analyzer** | LLM 意图识别与推理 | `llm` 服务（Prompt 工程 + 上下文管理）|
| **Synthesizer** | 结构化内容生成与输出 | `content generator` + SSE 推送 |
| **Scheduler** | 任务调度与状态管理 | `scheduler` 服务，支持定时/触发双模式 |

---

## 🛠️ Tech Stack & Performance

### 技术栈

| 层级 | 技术 |
|------|------|
| 后端 | Go 1.26, Gin, 自研微服务架构 |
| 前端 | Next.js 14 (App Router), TypeScript, Tailwind CSS |
| 状态管理 | Zustand |
| 实时通信 | SSE (Server-Sent Events) |
| AI 集成 | OpenAI API, Prompt 工程优化 |
| 部署 | Docker, Docker Compose |

### 性能指标

| 指标 | 数值 |
|------|------|
| SSE 首字节时间 (TTFB) | < 200ms |
| 端到端生成响应 | < 3s（标准查询）|
| API 限流保护 | 100 req/min |
| 容器启动时间 | < 5s |

---

## 📋 系统要求

- Docker 20.10+
- Docker Compose 2.0+
- Git
- Node.js 18+
- Go 1.26+

---

## 🚀 快速开始

### 1. 克隆项目
```bash
git clone https://github.com/haoyuehhh/opflow.git
cd opflow
```

### 2. 环境配置

#### 后端环境变量
创建 `backend/.env` 文件：
```env
PORT=8080
ENV=development
OPENAI_API_KEY=your_openai_api_key_here
```

#### 前端环境变量
创建 `frontend/.env.local` 文件：
```env
NEXT_PUBLIC_API_URL=http://localhost:8080
```

### 3. 启动服务
```bash
# 启动所有服务
docker-compose up --build

# 或分别启动
# 启动后端
cd backend
go run cmd/server/main.go

# 启动前端
cd frontend
npm run dev
```

### 4. 访问应用
- 前端应用: http://localhost:3000
- 后端 API: http://localhost:8080
- 健康检查: http://localhost:8080/health

---

## 📁 项目结构

```
opflow/
├── backend/                 # Go 后端服务
│   ├── cmd/                 # 命令行入口
│   ├── internal/            # 内部包
│   │   ├── config/          # 配置管理
│   │   ├── handlers/        # HTTP 处理器
│   │   ├── middleware/      # 中间件
│   │   ├── models/          # 数据模型
│   │   ├── pkg/             # 工具包
│   │   ├── services/        # 业务服务
│   │   │   ├── crawler/     # 采集服务 (Collector)
│   │   │   ├── llm/         # LLM 推理服务 (Analyzer)
│   │   │   ├── content/     # 内容合成服务 (Synthesizer)
│   │   │   ├── scheduler/   # 任务调度
│   │   │   └── sse/         # 实时推送服务
│   │   └── utils/           # 工具函数
│   └── go.mod               # Go 模块依赖
├── frontend/                # Next.js 前端应用
│   ├── app/                 # 页面和路由
│   ├── components/          # React 组件
│   ├── lib/                 # 库和工具
│   └── package.json         # Node.js 依赖
├── docker-compose.yml       # Docker 编排配置
└── README.md               # 项目文档
```

---

## 🔧 API & Effect Evaluation

### 核心接口
- `GET /api/v1/hotspots` — 多源热点采集
- `POST /api/v1/hotspots/:id/generate` — LLM 推理与内容合成
- `GET /health` — 服务健康检查

### 请求示例

#### 获取热点
```bash
curl http://localhost:8080/api/v1/hotspots
```

#### 生成内容
```bash
curl -X POST http://localhost:8080/api/v1/hotspots/zhihu_1/generate \
  -H "Content-Type: application/json" \
  -d '{"platform": "structured_output"}'
```

### 效果评估维度

| 维度 | 评估方法 |
|------|---------|
| 响应耗时 | SSE TTFB < 200ms，端到端 < 3s |
| 生成稳定性 | Prompt 模板版本管理 + 输出格式校验 |
| 服务可用性 | Docker 容器化 + 速率限制中间件保护 |
| 输出质量 | 结构化 JSON Schema 约束，前端可渲染 |

---

## 🤝 贡献指南

我们欢迎所有形式的贡献！请遵循以下步骤：

1. Fork 项目
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

### 开发规范
- 遵循 Go 的代码风格指南
- 编写单元测试
- 更新相关文档
- 确保代码通过所有测试

---

## 📝 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

---

## 📞 联系方式

- 项目维护者: haoyuehhh
- GitHub: https://github.com/haoyuehhh/opflow

---

*OpFlow — 让数据流转更智能，让 AI 合成更高效！*
