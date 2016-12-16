/**
 * Created by Administrator on 2016/12/12.
 */
var messageTxt;
var messages;

$(function () {

    messageTxt = $("#messageTxt");
    messages = $("#messages");


    ws = new Ws("ws://" + HOST + "/doubeles");
    ws.OnConnect(function () {
        console.log("Websocket connection enstablished");
    });

    ws.OnDisconnect(function () {
        appendMessage($("<div><center><h3>Disconnected</h3></center></div>"));
    });

    ws.On("chat", function (message) {
        console.log(message+"----");
        appendMessage($("<div>" + message + "</div>"));
    })

    $("#sendBtn").click(function () {
        //ws.EmitMessage(messageTxt.val());
        console.log(messageTxt.val().toString());
        ws.Emit("chat", messageTxt.val().toString());
        messageTxt.val("");
    })

})


function appendMessage(messageDiv) {
    var theDiv = messages[0]
    var doScroll = theDiv.scrollTop == theDiv.scrollHeight - theDiv.clientHeight;
    messageDiv.appendTo(messages)
    if (doScroll) {
        theDiv.scrollTop = theDiv.scrollHeight - theDiv.clientHeight;
    }
}