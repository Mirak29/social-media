<script>
	import axios from "axios";
	import { onMount } from "svelte";
	import { PUBLIC_BACKEND_URL_FRONTEND,PUBLIC_SOCKET_URL } from '$env/static/public'; 
	let NotifSocket;
	export let data;

	const notificationRequireConfirmation = (type) => {
		return ["follow", "join-Group", "invited-to-join-Group"].includes(type);
	};

	const getNotifMessageSwitchType = (type, name, group_title) => {
		let message = "";
		switch (type) {
			case "follow":
				message = name + " request to following you.";
				break;
			case "invited-to-join-Group":
				message = `${name} invite you to join #${group_title} group.`;
				break;
			case "join-Group":
				message = `${name} request to join ${group_title}  group.`;
				break;
			case "succesrequest":
				if (group_title != "") {
					message = `your request to join ${group_title} group is accepted.`;
				} else {
					message = `${name} accepted your follow request.`;
				}
				break;
			case "followInfo":
				message = `${name} is following you.`;
				break;
			case "event-notif":
				message = `${name} created a new event.`;
				break;
			default:
				message = "no message";
				break;
		}
		return message;
	};

	const markAsRead = async (notifId, receiverId) => {
		try {
			let header = {
				cookie: document.cookie,
			};

			const config = {
				method: "get",
				withCredentials: true,
				header,
				mode: "no-cors",
				params: { user_id: receiverId, notif_id: notifId },
			};

			let response = await axios(
				`${PUBLIC_BACKEND_URL_FRONTEND}/notifAsRead`,
				config
			);
			if (response?.data?.success) {
				data.notifications = data.notifications.filter(
					(notif) => notif.id !== notifId
				);
			}
		} catch (error) {
			console.log(error);
		}
	};

	const respondToRequest = async (
		accept,
		notifId,
		receiverId,
		group_id = 0
	) => {
		try {
			let header = {
				cookie: document.cookie,
			};

			const config = {
				method: "get",
				withCredentials: true,
				header,
				mode: "no-cors",
				params: { user_id: receiverId, notif_id: notifId, group_id, accept },
			};
			let response = await axios(
				`${PUBLIC_BACKEND_URL_FRONTEND}/notiftraitement`,
				config
			);
			if (response?.data?.success) {
				console.log('RESPONSE : ', response?.data);
				data.notifications = data.notifications.filter(
					(notif) => notif.id !== notifId
				);
				console.log(notifId);
			}
		} catch (error) {
			console.log(error);
		}
	};

	onMount(async () => {
		if (!NotifSocket) {
			NotifSocket = new WebSocket(PUBLIC_SOCKET_URL + "/initnotifsocket");
		}

		NotifSocket.onmessage = function (event) {
			let newEvent = JSON.parse(event.data);

			console.log('notifs : ',newEvent);

			data.notifications = [...data.notifications, newEvent.data];
		};

		NotifSocket.onclose = function (event) {
			NotifSocket = null;
		};

		NotifSocket.onerror = function (error) {
			NotifSocket = null;
		};
	});
</script>
<style>
	.notifications_network::-webkit-scrollbar{
		display: none;
	}
</style>

<a
	href="#"
	class="p-2 text-center ms-auto menu-icon"
	id="dropdownMenu3"
	data-bs-toggle="dropdown"
	aria-expanded="false"
>
	{#if data.notifications.length > 0}
		<span class="dot-count bg-warning"></span>
	{/if}
	<i class="feather-bell font-xl text-current"></i>
</a>
<div
	class="dropdown-menu dropdown-menu-end p-4 rounded-3 border-0 shadow-lg notifications_network" style="height: 500px; overflow-y:scroll;"
	
>

	<h4 class="fw-700 font-xss mb-4"  >Notifications</h4>
	{#if data}
		{#key data}
			{#each data.notifications as notif}
				<div class="card bg-transparent-card w-100 border-0 ps-5 mb-3">
					<img
						class="w40 position-absolute left-0"
						src="//ui-avatars.com/api/?name={notif.firstname +
							' ' +
							notif.lastname}&size=100&rounded=true&color=fff&background=random"
						width="50"
						alt={notif?.firstName + " " + notif?.lastName}
					/>
					<h5 class="font-xsss text-grey-900 mb-1 mt-0 fw-700 d-block">
						{notif?.firstname + " " + notif?.lastname}
					</h5>
					<h6 class="text-grey-500 fw-500 font-xssss lh-4">
						{getNotifMessageSwitchType(
							notif.type,
							notif?.firstname + " " + notif.lastname,
							notif.grouptitle
						)}
					</h6>
				</div>

				<div class="card-body d-flex align-items-center pt-0 ps-4 pe-4 pb-4">
					{#if notificationRequireConfirmation(notif.type)}
						<a
							href="#"
							on:click={() => {
								respondToRequest(1, notif.id, notif.sender_id, notif.group_id);
							}}
							class="p-2 lh-20 w100 bg-primary-gradiant me-2 text-white text-center font-xssss fw-600 ls-1 rounded-xl"
							>Confirm</a 
						>
						<a
							href="#"
							on:click={() => {
								respondToRequest(0, notif.id, notif.sender_id, notif.group_id);
							}}
							class="p-2 lh-20 w100 bg-grey text-grey-800 text-center font-xssss fw-600 ls-1 rounded-xl"
							>Decline</a
						>
					{:else}
						<a
							href="#"
							on:click={() => {
								markAsRead(notif.id, notif.sender_id);
							}}
							class="p-2 lh-20 w100 bg-grey text-grey-800 text-center font-xssss fw-600 ls-1 rounded-xl"
							>Delete</a
						>
					{/if}
				</div>
			{/each}
		{/key}
	{/if}
	<!-- {/if} -->
</div>
