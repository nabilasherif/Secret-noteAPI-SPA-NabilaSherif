import { createRouter, createWebHistory } from 'vue-router';
import Home from '../views/HomePage.vue';
import Profile from '../views/Profile.vue';
import PageNotFound from "../views/PageNotFound.vue";
import ViewNote from '../views/ViewNote.vue';

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/profile',
    name: 'Profile',
    component: Profile
  },
  {
    path: '/note/:note_url',
    name: 'ViewNote',
    component: ViewNote
  },
  {
    path: "/:catchAll(.*)",
    name: "notfound",
    component: PageNotFound,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
