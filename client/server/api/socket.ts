import type { Peer } from "crossws";
import { getQuery } from "ufo";

const notifconns = new Map<string, { conn: Peer, onLine: boolean }>();

export const notifUser = (data: any) => {
  const user = notifconns.get(data.concernID)
  if (user) {
    if (user?.onLine) {
      user.conn.send(data)
    }
  }
}

export default defineWebSocketHandler({
  async open(peer: Peer) {
    const query = get(peer)
    const channel = query.channel as string
    const userId = query.userId as string;
    if (channel === "notif") {
      notifconns.set(userId, { conn: peer, onLine: true });
    }
  },
  // async message(peer: Peer, message) {
  //   console.log(`[ws] message ${peer} ${message.text()}`);


  //   if (message.text() === "ping") {
  //     peer.send({ user: "server", message: "pong" });
  //     return
  //   }

  //   const _message = {
  //     user: "TEST",
  //     message: message.text(),
  //   };
  //   peer.send(_message); // echo back
  //   peer.publish("chat", _message);

  //   // Store message in database
  // },

  // close(peer: Peer, details) {
  //   console.log(`[ws] close ${peer}`);
  //   const query = get(peer)
  //   const channel = query.channel as string
  //   const userId = query.userId as string;
  //   if (channel === "notif") {
  //     notifconns.set(userId, { conn: peer, onLine: false });
  //   }

  //   // const userId = getUserId(peer) /;
  //   // users.set(userId, { online: false });
  // },

  // error(peer: Peer, error) {
  //   console.log(`[ws] error ${peer}`, error);
  // },

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
