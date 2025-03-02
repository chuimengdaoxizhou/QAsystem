import asyncio

from langchain.agents.chat.output_parser import ChatOutputParser
from langchain_core.prompts import ChatPromptTemplate

from langchain_ollama import ChatOllama

from text_generate_server.server import Server


def chatollama():
    llm = ChatOllama(model="qwen2.5:0.5b", base_url="localhost:11434")

    parser = ChatOutputParser()
    prompt = ChatPromptTemplate.from_template(
        """
        你是一个知识库，回答用户的问题或者总结
    
        用户输入的内容:{question}
        """
    )

    chain = prompt | llm

    question = "ollama是谁开发的"

    chunks = []

    for chunk in llm.stream(question):
        chunks.append(chunk)
        print(chunk.content, end="|", flush=True)

    return

if __name__ == "__main__":
    server = Server()
    server.serve()

