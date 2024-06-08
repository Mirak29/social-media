<script>
    import {onMount} from "svelte"
    import { page } from "$app/stores"
    import EmojiPicker from 'svelte-emoji-picker';
    import { afterUpdate } from 'svelte';
    import {LastMessage} from "../../stores"
    import {ContactsStore} from "../../stores"
    let Users
    let usersOnline
    ContactsStore.subscribe((val) => {Users = val})

    import {WS} from "../../socket"
    export let data
    $:console.log('data chat', data);
    let element;

    
    $:console.log('data GUEYE receiver', data);
    $:allmessage = data?.res?.result?.messages
    afterUpdate(() => {
            console.log("afterUpdate");
            if(allmessage) scrollToBottom(element);
    });
    $: if(allmessage && element) {
		console.log("tick");
		scrollToBottom(element);
	}
    let socket
    $:receiver_id = $page.params.id
    $:console.log('RECEIVER ID', receiver_id);
    let lastMesage = ''
    onMount(async () => { 
        let audioFile = new Audio("/audio/notif.mp3")
        WS.subscribe((val)=> socket = val)
        socket.addEventListener("open", ()=> {
            console.log("Opened")
        })
        socket.onmessage = (e) => {
           
            var newMessage = JSON.parse(e.data);
            if (newMessage.action == 'onlineUsers') {
                usersOnline = newMessage.users
                for (let user of Users) {
                    if (usersOnline.some(onlineUser => onlineUser.ID === user.ID)) {
                        user.Online = true
                    } else {
                        user.Online = false
                    }
                }   
                ContactsStore.set(Users)
            }
            if ((newMessage.sender_id == receiver_id || newMessage.receiver_id == receiver_id ) && newMessage.groupId == 0) {
             
                if (allmessage) {
                    allmessage = [...allmessage , newMessage]
                } else {
                    allmessage = [newMessage]
                }
            }
            if (newMessage.sender_id == receiver_id ) {
                audioFile.play()
                LastMessage.set(newMessage)
            }
            
            
        }
    })
    const scrollToBottom = async (node) => {
        node.scroll({ top: node.scrollHeight, behavior: 'smooth' });
    };
    let Message = ''
    
    let showEmojis = false;

    function SendMessage() {
        let mess = {
            receiver_id : Number(receiver_id),
            content : Message,
        }
        Message = ''
        const messageJSON = JSON.stringify(mess);
        socket.send(messageJSON)
    }
    function Send(e) {
        if (e.code == "Enter" ) {
            SendMessage()
        }
    }
</script>

<div class="main-content right-chat-active">
            
    <div class="middle-sidebar-bottom">
        <div class="middle-sidebar-left pe-0 ps-lg-3 ms-0 me-0" style="max-width: 100%;">
            <div class="row">
                   

                <div class="col-lg-12 position-relative">
                    <div class="chat-wrapper pt-0 w-100 position-relative scroll-bar bg-white theme-dark-bg" bind:this={element} style="overflow:auto;">
                        <div class="chat-body p-3 " >
                            <div class="modal-popup-header w-100 border-bottom text-center">
                                <div class="card p-3 d-block border-0 d-block d-flex justify-content-center align-items-center">
                                    <figure class="avatar mb-0 me-2">
                                        {#if data?.res?.result?.receiver?.Avatar}
                                            <img src={data?.res?.result?.receiver?.Avatar} alt="display photo" class="w35 me-1">
                                         {:else}
                                            <img src="//ui-avatars.com/api/?name={data?.res?.result?.receiver?.FirstName +' '+data?.res?.result?.receiver?.LastName}e&size=100&rounded=true&color=fff&background=random"
                                            alt="{data?.res?.result?.receiver?.FirstName +' ' + data?.res?.result?.receiver?.LastName}"
                                            class="shadow-sm rounded-circle w45"
                                                />
                                        {/if}
                                    </figure>
                                    <div>
                                        <h5 class="fw-700 text-black font-xssss mt-1 mb-1">{data?.res?.result?.receiver?.FirstName} {data?.res?.result?.receiver?.LastName}</h5>
                                    </div>
                                </div>
                            </div>
                            <div class="messages-content pb-5" >
                                {#if allmessage?.length > 0}
                                    {#each allmessage as message }
                                    {#if message.receiver_id != receiver_id}
                                        <div class="message-item" >
                                            <div class="message-user">
                                                <figure class="avatar">
                                                    {#if !message?.avatar}
                                                        <img src="//ui-avatars.com/api/?name={message?.firstName +' '+message?.lastName}e&size=100&rounded=true&color=fff&background=random"
                                                        alt="{message?.firstName +' '+message?.lastName}"
                                                        class="shadow-sm rounded-circle w45"
                                                        />
                                                    {:else}
                                                        <img src={message.avatar} alt="{message.receiver_id }">
                                                    {/if}
                                                </figure>
                                                <div>
                                                    <h5>{message.firstName} {message.lastName}</h5>
                                                    <div class="time">01:35 PM</div>
                                                </div>
                                            </div>
                                            <div class="message-wrap">{message.content}</div>
                                        </div>
                                    {:else}
                                        <div class="message-item outgoing-message">
                                            <div class="message-user">
                                                <figure class="avatar">
                                                    {#if !message?.avatar}
                                                        <img src="//ui-avatars.com/api/?name={message?.firstName +' '+message?.lastName}e&size=100&rounded=true&color=fff&background=random"
                                                        alt="{message?.firstName +' '+message?.lastName}"
                                                        class="shadow-sm rounded-circle w45"
                                                    />
                                                    {:else}
                                                       <img src={message.avatar} alt="display p{message.receiver_id }">
                                                    {/if}
                                                </figure>
                                                <div>
                                                    <h5>{message.firstName} {message.lastName}</h5>
                                                    <div class="time">01:35 PM<i class="ti-double-check text-info"></i></div>
                                                </div>
                                            </div>
                                            <div class="message-wrap">{message.content}</div>
                                        </div>
                                    {/if} 
                                   
                                    {/each}
                                {/if}
                                <div class="clearfix"></div>
                            </div>
                        </div>
                    </div>
                    <div class="chat-bottom dark-bg p-3 shadow-none theme-dark-bg" style="width: 98%;">
                        <div class="chat-form">
                            <div style="display: {showEmojis ? 'block' : 'none'};" >
                                <EmojiPicker bind:value={Message} />
                            </div>

                            
                            <div class="form d-flex w-80 gap-1">
                                <button class="bg-current float-left" on:click={() => showEmojis = !showEmojis}>😀</button>
                                <input type="text" placeholder="Start typing.." class="border"  name="name"   bind:value={Message} style="color: black;" on:keydown={Send}>
                                <button class="bg-current" on:click={SendMessage}><i class="ti-arrow-right text-white" ></i></button>
                            </div>          
                        </div>
                    </div> 
                </div>

            </div>
        </div>
         
    </div>            
</div>