import grpc
from concurrent import futures
import asyncio
from langchain_core.output_parsers import StrOutputParser
from langchain_ollama import ChatOllama
from langchain_core.prompts import ChatPromptTemplate

from text_generate_server.proto import server_pb2, server_pb2_grpc


class Server(server_pb2_grpc.ModelServiceServicer):
    def __init__(self):
        pass

    async def async_stream(self, chain, question):
        # 异步地遍历stream返回的token
        async for message in chain.astream({"question": question}):
            yield message

    async def GenerateContentStream(self, request, context):
        llm = ChatOllama(model="qwen2.5:0.5b", base_url="http://localhost:11434")
        parser = StrOutputParser()

        prompt = ChatPromptTemplate.from_template(
            """
            你是一个全能知识库，详细的回答用户的问题

            用户的问题是：{question}
            """
        )

        question = request.prompt
        chain = prompt | llm | parser

        # 异步生成器
        async for token in self.async_stream(chain, question):
            yield server_pb2.Response(token=token)


async def serve():
    # 使用异步gRPC服务器
    server = grpc.aio.server()
    server_pb2_grpc.add_ModelServiceServicer_to_server(Server(), server)

    server.add_insecure_port("[::]:50051")
    print("Server is running on port 50051...")
    await server.start()
    await server.wait_for_termination()


if __name__ == "__main__":
    asyncio.run(serve())
