import Vue from "vue";
import VueRouter from "vue-router";
import Note from "@/components/Note.vue";
import 'bootstrap/dist/css/bootstrap.css';
import CreateUser from"@/components/CreateUser.vue";

Vue.useAttrs(VueRouter);
///rest of url\\
const routes = [
    {
        path:'/note',
        name: 'Note',
        component : Note,
    },
    {
        path: '/createuser',
        name :'CreateUser',
        component : CreateUser,
    }
];


const router = new VueRouter({
    mode: "history",
    base: process.env.BASE_URL,
    routes,
});

export default router;