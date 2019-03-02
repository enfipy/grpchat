import * as grpcWeb from 'grpc-web';
import {
  ClientMessage,
  GetMessagesRequest,
  SendMessageResponse,
  ServerMessage} from './chat_pb';

export class ChatClient {
  constructor (hostname: string,
               credentials: null | { [index: string]: string; },
               options: null | { [index: string]: string; });

  getMessages(
    request: GetMessagesRequest,
    metadata: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<ServerMessage>;

  sendMessage(
    request: ClientMessage,
    metadata: grpcWeb.Metadata,
    callback: (err: grpcWeb.Error,
               response: SendMessageResponse) => void
  ): grpcWeb.ClientReadableStream<SendMessageResponse>;

}

export class ChatPromiseClient {
  constructor (hostname: string,
               credentials: null | { [index: string]: string; },
               options: null | { [index: string]: string; });

  getMessages(
    request: GetMessagesRequest,
    metadata: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<ServerMessage>;

  sendMessage(
    request: ClientMessage,
    metadata: grpcWeb.Metadata
  ): Promise<SendMessageResponse>;

}

