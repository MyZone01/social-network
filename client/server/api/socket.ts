import type { Peer } from "crossws";
import { getQuery } from "ufo";
import { server_socket } from "~/stores/socketCon";

const conns = new Map<string, { conn: Peer }>();

export default defineWebSocketHandler({
  async open(peer: Peer) {
    console.log(`[ws] open ${peer}`);
    const query = get(peer)
    const channel = query.channel as string
    const userId = query.userId as string;
    // users.set(userId, { online: true });
    console.log(channel, userId);

    if (channel === "notif") {
      conns.set(userId, { conn: peer });
      server_socket!.onmessage = (event) => {
        console.log(event);
        // peer.send(event.data);
      };
    }
  },
  async message(peer: Peer, message) {
    console.log(`[ws] message ${peer} ${message.text()}`);


    if (message.text() === "ping") {
      peer.send({ user: "server", message: "pong" });
      return
    }

    const _message = {
      user: "TEST",
      message: message.text(),
    };
    peer.send(_message); // echo back
    peer.publish("chat", _message);

    // Store message in database
  },

  close(peer: Peer, details) {
    console.log(`[ws] close ${peer}`);

    // const userId = getUserId(peer) /;
    // users.set(userId, { online: false });
  },

  error(peer: Peer, error) {
    console.log(`[ws] error ${peer}`, error);
  },

  upgrade(req) {
    return {
      headers: {
        "x-powered-by": "cross-ws",
      },
    };
  },
});

function get(peer: Peer) {
  const query = getQuery(peer.url);
  return query
}

// function getStats() {
//   const online = Array.from(users.values()).filter((u) => u.online).length;
//   return { online, total: users.size };
// }
