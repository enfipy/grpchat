export class ClientMessage {
  constructor ();
  getMessage(): string;
  setMessage(a: string): void;
  toObject(): ClientMessage.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => ClientMessage;
}

export namespace ClientMessage {
  export type AsObject = {
    Message: string;
  }
}

export class GetMessagesRequest {
  constructor ();
  getUsername(): string;
  setUsername(a: string): void;
  toObject(): GetMessagesRequest.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => GetMessagesRequest;
}

export namespace GetMessagesRequest {
  export type AsObject = {
    Username: string;
  }
}

export class SendMessageResponse {
  constructor ();
  toObject(): SendMessageResponse.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => SendMessageResponse;
}

export namespace SendMessageResponse {
  export type AsObject = {
  }
}

export class ServerMessage {
  constructor ();
  getUsername(): string;
  setUsername(a: string): void;
  getMessage(): string;
  setMessage(a: string): void;
  getCreatedAt(): number;
  setCreatedAt(a: number): void;
  toObject(): ServerMessage.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => ServerMessage;
}

export namespace ServerMessage {
  export type AsObject = {
    Username: string;
    Message: string;
    CreatedAt: number;
  }
}

