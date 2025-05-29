from langchain_openai import ChatOpenAI

llm = ChatOpenAI(
  model="qwen-plus",
  api_key="sk-2bb881d94f964955af76d6523c127218",
  base_url="https://dashscope.aliyuncs.com/compatible-mode/v1",
)
ret = llm.invoke("Hello, world!")

print(ret)