import type { Peer } from "crossws";
import { getQuery } from "ufo";
import { server_socket } from "~/stores/socketCon"

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
  async message(peer: Peer, message) {
    const data = JSON.parse(message.toString())
    if (data.type && data.id) {
      server_socket!.send(data)
    }
  },

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
