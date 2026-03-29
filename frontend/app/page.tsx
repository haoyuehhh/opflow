'use client'

import { useEffect, useState } from 'react'

export default function Home() {
  const [hotspots, setHotspots] = useState<any[]>([])
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(true)
  const [generatingId, setGeneratingId] = useState<string | null>(null)
  const [generatedContent, setGeneratedContent] = useState('')
  const [showModal, setShowModal] = useState(false)
  const [selectedPlatform, setSelectedPlatform] = useState('xiaohongshu')

  useEffect(() => {
    fetch('http://localhost:8080/api/v1/hotspots')
      .then(res => {
        if (!res.ok) throw new Error('Network response was not ok')
        return res.json()
      })
      .then(data => {
        setHotspots(data.data || [])
        setLoading(false)
      })
      .catch(err => {
        setError('Failed to load: ' + err.message)
        setLoading(false)
      })
  }, [])

  const handleGenerate = async (id: string, title: string) => {
    setGeneratingId(id)
    
    // 前端模板方案（模拟AI生成延迟）
    setTimeout(() => {
      const templates: Record<string, string> = {
        xiaohongshu: `🔥${title}\n\n家人们谁懂啊！${title}真的太绝了！\n\n今天必须跟大家聊聊这个话题，全是干货...\n\n💡核心观点：\n✅ 第一，这个现象反映了行业新趋势\n✅ 第二，值得关注的变化正在发生\n✅ 第三，普通人如何抓住机会\n\n看完这篇你就懂了！建议收藏👇\n\n#${title.slice(0,4)} #热点话题 #AI运营 #深度解析 #必看`,

        weibo: `【${title}】\n\n${title}引发全网热议！🔥\n\n这件事的核心在于：当技术与需求碰撞，产生的化学反应远超预期。\n\n几个观察角度：\n1️⃣ 行业层面：这是趋势还是泡沫？\n2️⃣ 用户层面：真实需求被满足了吗？\n3️⃣ 未来走向：短期炒作还是长期价值？\n\n你怎么看？评论区聊聊👇\n\n转发关注，了解更多资讯\n\n#${title}##AI发展趋势##行业观察#`,

        gongzhonghao: `标题：深度解析：${title}背后的逻辑与机遇\n\n正文：\n\n近日，${title}引发广泛关注。作为从业者，我们认为这个现象值得深入剖析。\n\n一、现象回顾\n${title}的爆发并非偶然，而是多重因素叠加的结果。从市场环境看...；从技术演进看...；从用户行为看...\n\n二、核心逻辑\n1. 技术成熟度到达临界点\n2. 市场需求被精准击中\n3. 资本与流量的双重助推\n\n三、未来展望\n短期看热度，长期看价值。${title}能否持续，关键在于是否真正解决了痛点问题。\n\n结语：\n每个热点背后都有值得学习的逻辑。关注OpFlow，带你穿透表象看本质。\n\n（点击下方阅读原文，查看完整分析报告）`
      }
      
      const content = templates[selectedPlatform] || templates.xiaohongshu
      setGeneratedContent(content)
      setShowModal(true)
      setGeneratingId(null)
    }, 800) // 模拟800ms生成延迟，演示更真实
  }

  if (loading) return <div className="p-8 text-center">加载中...</div>
  if (error) return <div className="p-8 text-red-600">{error}</div>

  return (
    <div className="min-h-screen bg-gray-50 p-8">
      <div className="max-w-4xl mx-auto">
        <div className="flex justify-between items-center mb-6">
          <h1 className="text-3xl font-bold text-gray-800">OpFlow 热点监控</h1>
          <div className="flex items-center gap-2">
            <label className="text-sm text-gray-600">发布平台:</label>
            <select 
              value={selectedPlatform} 
              onChange={(e) => setSelectedPlatform(e.target.value)}
              className="border rounded px-2 py-1 text-sm"
            >
              <option value="xiaohongshu">小红书</option>
              <option value="weibo">微博</option>
              <option value="gongzhonghao">公众号</option>
            </select>
          </div>
        </div>
        
        <div className="grid gap-4">
          {hotspots.map((h) => (
            <div key={h.id} className="bg-white rounded-lg shadow p-4 hover:shadow-md transition">
              <div className="flex justify-between items-start">
                <h3 className="font-bold text-lg text-gray-800 flex-1">{h.title}</h3>
                <span className={`px-2 py-1 rounded text-sm ${
                  h.source === 'zhihu' 
                    ? 'bg-blue-100 text-blue-800' 
                    : 'bg-red-100 text-red-800'
                }`}>
                  {h.source}
                </span>
              </div>
              <div className="mt-3 flex justify-between items-center">
                <div className="text-sm text-gray-600">
                  热度: {h.score} | 
                  <a href={h.url} target="_blank" className="text-blue-600 hover:underline ml-2">查看原文</a>
                </div>
                <button
                  onClick={() => handleGenerate(h.id, h.title)}
                  disabled={generatingId === h.id}
                  className={`px-4 py-2 rounded text-sm font-medium ${
                    generatingId === h.id
                      ? 'bg-gray-300 text-gray-500 cursor-not-allowed'
                      : 'bg-blue-600 text-white hover:bg-blue-700'
                  }`}
                >
                  {generatingId === h.id ? '生成中...' : '生成文案'}
                </button>
              </div>
            </div>
          ))}
        </div>
      </div>

      {/* Modal */}
      {showModal && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4 z-50">
          <div className="bg-white rounded-lg max-w-lg w-full p-6">
            <div className="flex justify-between items-center mb-4">
              <h3 className="text-lg font-bold">AI生成文案</h3>
              <button onClick={() => setShowModal(false)} className="text-gray-500 hover:text-gray-700">✕</button>
            </div>
            <textarea
              value={generatedContent}
              readOnly
              className="w-full h-64 p-3 border rounded-lg resize-none text-sm"
            />
            <div className="mt-4 flex justify-end gap-2">
              <button 
                onClick={() => setShowModal(false)} 
                className="px-4 py-2 border rounded hover:bg-gray-50"
              >
                关闭
              </button>
              <button 
                onClick={() => navigator.clipboard.writeText(generatedContent)}
                className="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700"
              >
                复制文案
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  )
}