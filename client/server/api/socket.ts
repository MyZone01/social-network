import type { Peer } from "crossws";
import { getQuery } from "ufo";
import { server_socket } from "~/stores/socketCon"

const notifconns = new Map<string, { conn: Peer, onLine: boolean }>();
const messageconns = new Map<string, { conn: Peer, otherId: string }>();

export const notifUser = (data: any) => {
  const user = notifconns.get(data.concernID)
  if (user) {
    if (user?.onLine) {
      if (data.type === 'new_message') {
        const otherId = messageconns.get(data.concernID)?.otherId
        console.log(otherId, data.user.ID,"online message");
        if (!otherId || otherId !== data.user.ID) {
          user.conn.send(data)
        }
      } else {
        user.conn.send(data)
      }
    }
  }
}

export const messageUser = (data: any) => {
  console.log('messageUser', data);

  const user = messageconns.get(data.ReceiverID)
  if (user) {
    if (user.otherId === data.SenderID) {
      const online = notifconns.get(data.SenderID)?.onLine
      if (online) {
        user.conn.send({ online: true, ...data })
      } else {
        user.conn.send({ online: false, ...data })
      }
    } //else {
    //   const online = notifconns.get(data.SenderID)?.onLine
    //   if (online) {
    //     notifconns.get(data.SenderID)?.conn.send({ online: true, ...data })
    //   } else {
    //     notifconns.get(data.SenderID)?.conn.send({ online: false, ...data })
    //   }
    // }
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
    if (channel === "message") {
      const otherId = query.otherId as string;
      messageconns.set(userId, { conn: peer, otherId });
    }
  },
  async message(peer: Peer, message) {
    const data = JSON.parse(message.toString())
    const query = get(peer)
    const channel = query.channel as string
    const userId = query.userId as string;
    // if (channel === "notif") {
    //   notifUser(data)
    // }
    if (channel === "message") {
      server_socket!.send(data)
    }
  },

  close(peer: Peer, details) {
    console.log(`[ws] close ${peer}`);
    const query = get(peer)
    const channel = query.channel as string
    const userId = query.userId as string;
    if (channel === "notif") {
      notifconns.set(userId, { conn: peer, onLine: false });
    } else if (channel === "message") {
      messageconns.delete(userId);
    }



    // const userId = getUserId(peer) /;
    // users.set(userId, { online: false });
  },

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
