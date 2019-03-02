/**
 * @fileoverview gRPC-Web generated client stub for chat
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.chat = require('./chat_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.chat.ChatClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.chat.ChatPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!proto.chat.ChatClient} The delegate callback based client
   */
  this.delegateClient_ = new proto.chat.ChatClient(
      hostname, credentials, options);

};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.chat.GetMessagesRequest,
 *   !proto.chat.ServerMessage>}
 */
const methodInfo_Chat_GetMessages = new grpc.web.AbstractClientBase.MethodInfo(
  proto.chat.ServerMessage,
  /** @param {!proto.chat.GetMessagesRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.chat.ServerMessage.deserializeBinary
);


/**
 * @param {!proto.chat.GetMessagesRequest} request The request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.chat.ServerMessage>}
 *     The XHR Node Readable Stream
 */
proto.chat.ChatClient.prototype.getMessages =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/chat.Chat/GetMessages',
      request,
      metadata,
      methodInfo_Chat_GetMessages);
};


/**
 * @param {!proto.chat.GetMessagesRequest} request The request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.chat.ServerMessage>}
 *     The XHR Node Readable Stream
 */
proto.chat.ChatPromiseClient.prototype.getMessages =
    function(request, metadata) {
  return this.delegateClient_.client_.serverStreaming(this.delegateClient_.hostname_ +
      '/chat.Chat/GetMessages',
      request,
      metadata,
      methodInfo_Chat_GetMessages);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.chat.ClientMessage,
 *   !proto.chat.SendMessageResponse>}
 */
const methodInfo_Chat_SendMessage = new grpc.web.AbstractClientBase.MethodInfo(
  proto.chat.SendMessageResponse,
  /** @param {!proto.chat.ClientMessage} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.chat.SendMessageResponse.deserializeBinary
);


/**
 * @param {!proto.chat.ClientMessage} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.chat.SendMessageResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.chat.SendMessageResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.chat.ChatClient.prototype.sendMessage =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/chat.Chat/SendMessage',
      request,
      metadata,
      methodInfo_Chat_SendMessage,
      callback);
};


/**
 * @param {!proto.chat.ClientMessage} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.chat.SendMessageResponse>}
 *     The XHR Node Readable Stream
 */
proto.chat.ChatPromiseClient.prototype.sendMessage =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.sendMessage(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


module.exports = proto.chat;

