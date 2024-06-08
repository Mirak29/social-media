<script>
    import { onMount } from "svelte";
    import { page } from "$app/stores";
    import { WS } from "../../../socket";
    import EmojiPicker from "svelte-emoji-picker";
    import { afterUpdate } from "svelte";
    import { LastMessage } from "../../../stores";
    import { ContactsStore } from "../../../stores";
    export let data;
    $: console.log("DATA TO CHAT GROUP : ", data);
    let element;
    let showEmojis = false;

    let Users;
    let usersOnline;
    ContactsStore.subscribe((val) => {
        Users = val;
    });

    $: groupid = $page.params.id;
    $: allmessage = data?.res?.result?.messages;
    afterUpdate(() => {
        if (allmessage) scrollToBottom(element);
    });
    $: if (allmessage && element) {
        scrollToBottom(element);
    }
    $: currUserId = data?.res?.result?.userid;
    let socket;
    onMount(async () => {
        let audioFile = new Audio("/audio/notif.mp3");
        WS.subscribe((val) => (socket = val));
        socket.addEventListener("open", () => {
            console.log("Opened");
        });
        socket.onmessage = (e) => {
            var newMessage = JSON.parse(e.data);
            if (newMessage.action == "onlineUsers") {
                usersOnline = newMessage.users;
                for (let user of Users) {
                    if (
                        usersOnline.some(
                            (onlineUser) => onlineUser.ID === user.ID,
                        )
                    ) {
                        user.Online = true;
                    } else {
                        user.Online = false;
                    }
                }
                ContactsStore.set(Users);
            }
            if (newMessage?.groupId == groupid) {
                if (allmessage) {
                    allmessage = [...allmessage, newMessage];
                } else {
                    allmessage = [newMessage];
                }
            }
            if (newMessage.sender_id !== currUserId) {
                audioFile.play();
                LastMessage.set(newMessage);
            }
        };
    });
    const scrollToBottom = async (node) => {
        node.scroll({ top: node.scrollHeight, behavior: "smooth" });
    };
    let Message = "";
    function SendMessage() {
        if (Message != "") {
            let mess = {
                groupId: Number(groupid),
                content: Message,
            };
            Message = "";
            const messageJSON = JSON.stringify(mess);
            socket.send(messageJSON);
            showEmojis = false;
        }
    }
    function Send(e) {
        if (e.code == "Enter") {
            SendMessage();
        }
    }

    let count = 0; 
    let currenmember=data?.res?.result?.members .slice(0,4) || []
    function incrementCount() {
        count += 1;
    }
    
</script>

<div class="notifications"></div>
<div class="main-content right-chat-active">
    <div class="middle-sidebar-bottom">
        <div
            class="middle-sidebar-left pe-0 ps-lg-3 ms-0 me-0"
            style="max-width: 100%;"
        >
            <div class="row">
                <div class="col-lg-12 position-relative">
                    <div
                        class="chat-wrapper pt-0 w-100 position-relative scroll-bar bg-white theme-dark-bg"
                        bind:this={element}
                        style="overflow:auto;"
                    >
                        <div class="chat-body p-3">
                            <div
                                class="modal-popup-header w-100 border-bottom text-center"
                            >
                                <div
                                    class="card p-3 d-block border-0 d-block d-flex justify-content-center align-items-center"
                                >
                                    <figure class="avatar mb-0 me-2">
                                        <ul
                                            class="memberlist mt-1 mb-2 ms-0 d-block"
                                        >
                                        {#each currenmember as member}
                                            <li class="w20">
                                                {#if member.Avatar}
                                                        <img
                                                            src={member.Avatar}
                                                            alt="user"
                                                            class="w35 d-inline-block"
                                                            style="opacity: 1;"
                                                        />
                                                {:else}
                                                    <img src="//ui-avatars.com/api/?name={member.FirstName +' '+member.LastName}e&size=100&rounded=true&color=fff&background=random"
                                                    alt="{member.FirstName +' ' + member.LastName}"
                                                    class="w35 d-inline-block"  style="opacity: 1;"
                                                        />
                                                {/if}
                                            </li>
                                        {/each}
                                            <li class="last-member">
                                                <a
                                                    href=""
                                                    class="bg-greylight fw-600 text-grey-500 font-xssss w35 ls-3 text-center"
                                                    style="height: 35px; line-height: 35px;"
                                                    >+{ data?.res?.result?.members.length- currenmember.length }</a
                                                >
                                            </li>
                                            <li class="ps-3 w-auto ms-1">
                                                <a
                                                    href=""
                                                    class="fw-600 text-grey-500 font-xssss"
                                                    >Member(s)</a
                                                >
                                            </li>
                                        </ul>
                                    </figure>
                                    <div>
                                        <h5
                                            class="fw-700 text-black font-xssss mt-1 mb-1"
                                        >
                                            {data?.res?.result?.currentgroup
                                                ?.title}
                                        </h5>
                                    </div>
                                </div>
                            </div>
                            <div class="messages-content pb-5">
                                {#if allmessage?.length > 0}
                                    {#each allmessage as message}
                                        {#if message.sender_id != currUserId}
                                            <div class="message-item">
                                                <div class="message-user">
                                                    <figure class="avatar">
                                                        {#if !message?.avatar}
                                                            <img
                                                                src="//ui-avatars.com/api/?name={message?.firstName +
                                                                    ' ' +
                                                                    message?.lastName}e&size=100&rounded=true&color=fff&background=random"
                                                                alt={message?.firstName +
                                                                    " " +
                                                                    message?.lastName}
                                                                class="shadow-sm rounded-circle w45"
                                                            />
                                                        {:else}
                                                            <img
                                                                src={message?.avatar}
                                                                alt="display photo"
                                                            />
                                                        {/if}
                                                    </figure>
                                                    <div>
                                                        <h5>
                                                            {message.firstName}
                                                            {message.lastName}
                                                        </h5>
                                                        <div class="time">
                                                            01:35 PM
                                                        </div>
                                                    </div>
                                                </div>
                                                <div class="message-wrap">
                                                    {message.content}
                                                </div>
                                            </div>
                                        {:else}
                                            <div
                                                class="message-item outgoing-message"
                                            >
                                                <div class="message-user">
                                                    <figure class="avatar">
                                                        {#if !message?.avatar}
                                                            <img
                                                                src="//ui-avatars.com/api/?name={message?.firstName +
                                                                    ' ' +
                                                                    message?.lastName}e&size=100&rounded=true&color=fff&background=random"
                                                                alt={message?.firstName +
                                                                    " " +
                                                                    message?.lastName}
                                                                class="shadow-sm rounded-circle w45"
                                                            />
                                                        {:else}
                                                            <img
                                                                src={message?.avatar}
                                                                alt="display photo"
                                                            />
                                                        {/if}
                                                    </figure>
                                                    <div>
                                                        <h5>
                                                            {message.firstName}
                                                            {message.lastName}
                                                        </h5>
                                                        <div class="time">
                                                            01:35 PM<i
                                                                class="ti-double-check text-info"
                                                            ></i>
                                                        </div>
                                                    </div>
                                                </div>
                                                <div class="message-wrap">
                                                    {message.content}
                                                </div>
                                            </div>
                                        {/if}
                                    {/each}
                                {/if}
                                <div class="clearfix"></div>
                            </div>
                        </div>
                    </div>
                    <div
                        class="chat-bottom dark-bg p-3 shadow-none theme-dark-bg"
                        style="width: 98%;"
                    >
                        <div class="chat-form">
                            <div
                                style="display: {showEmojis
                                    ? 'block'
                                    : 'none'};"
                            >
                                <EmojiPicker bind:value={Message} />
                            </div>

                            <button
                                class="bg-current float-left"
                                on:click={() => (showEmojis = !showEmojis)}
                                >😀</button
                            >

                            <div class="form-group">
                                <input
                                    type="text"
                                    placeholder="Start typing.."
                                    bind:value={Message}
                                    style="color: black;"
                                    on:keydown={Send}
                                />
                            </div>
                            <button class="bg-current" on:click={SendMessage}
                                ><i class="ti-arrow-right text-white"
                                ></i></button
                            >
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
