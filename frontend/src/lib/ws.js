function joinRoom(roomId, userInfos) {
    const domainWithPort = `${window.location.hostname}:${window.location.port}`;
    const wsUrl = `ws://${domainWithPort}/joinRoom/${roomId}?userId=${userInfos.id}&username=${userInfos.username}`;
    const ws = new WebSocket(wsUrl);
    return ws
}

