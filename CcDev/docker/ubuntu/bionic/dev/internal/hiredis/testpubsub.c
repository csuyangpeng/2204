#include "epcredis.h"

void onMessage(redisAsyncContext *c, void *reply, void *privdata) {
    redisReply *r = reply;
    if (reply == NULL) return;

    if (r->type == REDIS_REPLY_ARRAY) {
        for (int j = 0; j < r->elements; j++) {
            printf("%u) %s\n", j, r->element[j]->str);
        }
    }
}

void onReply(redisAsyncContext *c, void *reply, void *privdata) {
    redisReply *r = reply;
    if (reply == NULL) return;

    if (r->type == REDIS_REPLY_ARRAY) {
        for (int j = 0; j < r->elements; j++) {
            printf("xxxxxx %u) %s\n", j, r->element[j]->str);
        }
    }
}

int main (int argc, char **argv) {
    /* print the version */
    EpcCtlMsg msg;
    msg.MsgType = 1;
    strcpy(msg.Info.UeInfo, "xxxUeInfo");
    msg.Info.Enb.EnbTeid = 123;
    strcpy(msg.Info.Enb.IpAddr, "192.138.12.3");
    msg.Info.GtpuNode.UpfTeid = 456;
    strcpy(msg.Info.GtpuNode.IpAddr, "192.138.12.2");
    strcpy(msg.Info.Apn.Name, "lily");
    strcpy(msg.Info.Apn.Ip, "10.18.1.22");

    printf("Version: %s\n", cJSON_Version());
    
    msg.MsgType = 1;
    sendtoepcupf((void*)&msg);

    return 0;
}
