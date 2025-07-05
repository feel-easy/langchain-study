from langgraph.prebuilt import create_react_agent
from langchain.chat_models import init_chat_model
import os
from pydantic import BaseModel, Field
from langchain_core.tools import tool

class MultiplyInputSchema(BaseModel):
    """Multiply two numbers"""
    a: int = Field(description="First operand")
    b: int = Field(description="Second operand")

@tool("multiply_tool", args_schema=MultiplyInputSchema)
def multiply(a: int, b: int) -> int:
    return a * b

os.environ["GOOGLE_API_KEY"] = "AIzaSyCBnVAKfKKaHqEeZAVw2afnCV3wY5Q38Vo"

llm = init_chat_model("google_genai:gemini-2.0-flash")
def get_weather(city: str) -> str:  
    """Get weather for a given city."""
    return f"It's always sunny in {city}!"

agent = create_react_agent(
    model=llm,  
    tools=[get_weather],  
    prompt="You are a helpful assistant"  
)
config = {"configurable": {"thread_id": "1"}}

# Run the agent
for stream_mode, chunk in agent.stream(
    {"messages": [{"role": "user", "content": "what is the weather in sf"}]},
    stream_mode=["updates", "messages", "custom"]
):
    print(chunk)
    print("\n")