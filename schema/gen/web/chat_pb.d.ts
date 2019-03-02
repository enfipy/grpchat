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

