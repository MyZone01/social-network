import { getNotifications } from "./getNotifications";



export const useClearNotif = async (id: string, action: string = "clear") => {
    const data = { type: action, id: id };
    const response = await fetch("/api/notification/clearnotification", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    }).then(async (res) => await res.json()).catch((err) => {
        return {
            status: 500,
            body: "Internal server error",
        };
    });
    return response;
}

export const notifications = useState<Notif[]>(() => []);

export type Notif = {
    id: number;
    notifId: string;
    type: string;
    message: string;
    user: any;
    created_at: Date;
};

export const addNotif = (notif: any) => {
    removeLasNotifType(notif)
    notifications.value.push({
        id: notifications.value.length + 1,
        notifId: notif.id,
        type: notif.type,
        message: notif.message,
        user: notif.user,
        created_at: new Date(notif.created_at),
    });
}

const removeLasNotifType = (notif: any) => {
    notifications.value.forEach((n, index) => {
        if (((notif.type === "follow_request"
            || notif.type === "follow_accepted"
            || notif.type === "follow_declined"
            || notif.type === "unfollow")
            && (
                n.type === "follow_request"
                || n.type === "follow_accepted"
                || n.type === "follow_declined"
                || n.type === "unfollow"))
            && notif.user.id === n.user.id) {
            notifications.value.splice(index, 1)
        }
    })
}

export const deleteNotif = (id: string) => {
    notifications.value.forEach((n, index) => {
        if (n.notifId == id) {
            notifications.value.splice(index, 1)
        }
    })
}

export const clearNotif = async (notif: Notif | undefined, action: string, message?: string) => {
    let type = 'clear'
    let notifId = notif?.notifId || ''
    if (action === 'all') {
        type = 'clear_all'
    }
    const res = await useClearNotif(notifId, type);

    if (action === 'redirect') {
        deleteNotif(notif!.notifId)
        navigateTo("/profile/" + notif!.user.nickname);
    } else if (action === 'all' && notifications.value.length > 0) {
        notifications.value = []
        notifications.value.push(
            {
                id: 1,
                notifId: '',
                type: "clear",
                message: "All notifications have been cleared",
                user: 'accepted',
                created_at: new Date()
            }
        )
        setTimeout(() => {
            notifications.value = []
        }, 5000);
    } else if (res.message) {
        notifications.value[notif!.id - 1].message = message!
        notifications.value[notif!.id - 1].type = "clear"
        setTimeout(() => {
            deleteNotif(notif!.notifId)
        }, 5000);
    }
}

if (!notifications.value.length) {
    const notifs = await getNotifications();
    if (!notifs.error && Array.isArray(notifs)) {
        Array.from(notifs).forEach((notif: any) => {
            addNotif(notif)
        })
        notifications.value.sort((a, b) => b.created_at.getTime() - a.created_at.getTime());
    }
}