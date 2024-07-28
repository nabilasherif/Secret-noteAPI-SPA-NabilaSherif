<template>
    <v-container>
        <v-row class="justify-end">
            <v-btn color="info" @click="logOut">
                <img src="/root/Secret-noteAPI-SPA-NabilaSherif/frontend/src/assets/logout.png" width=20px>
                Log Out
            </v-btn>
        </v-row>
        <v-row style="padding-left: 5cm;">
            <v-col>
                <v-card max-width="800px" >
                    <v-card-title>POST NEW NOTE</v-card-title>
                    <v-form @submit.prevent="postNote" >
                        <v-card-subtitle>
                            <v-text-field v-model="createNewNote.notetext" label="Note Text" />
                            <v-text-field v-model="createNewNote.expirationdate" label="Expiration Date" type="date" />
                            <v-text-field v-model="createNewNote.maxviews" label="Max Views" type="number" />
                            <v-card-actions>
                            <v-spacer></v-spacer>
                            <v-btn type="submit" color="primary">POST</v-btn>
                            </v-card-actions> 
                        </v-card-subtitle>
                    </v-form>
                </v-card>   
            </v-col>
        </v-row>
        <br><hr><br>
        <v-row style="padding-left: 1cm;">
            <h2 class="blue--text text--darken-2 display-2 font-weight-bold ">MY NOTES 📝</h2>
            <div v-for=" note in userNotes" :key="note.Url" >
                <p>
                    <strong>Note Text:</strong> {{ note.NoteText }}<br>
                    <strong>Current Views:</strong> {{ note.CurrentViews }}<br>
                    <strong>Max Views:</strong> {{ note.MaxViews }}<br>
                    <strong>Expiration Date:</strong> {{ note.ExpirationDate }}<br>
                    <strong>URL:</strong> <a :href="`http://localhost:5173/note/${note.Url}`" target="_blank">http://localhost:5173/note/{{ note.Url }}</a>
                </p>
                <hr>
            </div>
        </v-row>
    </v-container>
</template>

<script setup>
import {ref, onMounted} from 'vue';
import { useToast } from 'vue-toastification';
import userService from '../services/userservice.ts';
import { useRouter } from 'vue-router';

const toast = useToast();
const router = useRouter();

const showCreate=ref(true);

const createNewNote=ref( {
    notetext: "",
    expirationdate: "",
    maxviews: "",
});

const userNotes=ref([])

function initForm() {
    createNewNote.value.notetext = "";
    createNewNote.value.expirationdate = "";
    createNewNote.value.maxviews = "";
}

function onSubmitPost(){
    postNote()
    initForm()
}

async function postNote() {
    await userService.authClient().post("/note", createNewNote.value)
    .then(response => {
            console.log(response.data);
            toast.success('Note created successfully');
            initForm();
    })
    .catch(err => {
        toast.error('Error creating note');
        console.error('Error creating note:', err);
    });
}

function  getMyNotes() {
    userService.baseClient().get("/notes")
    .then(res => {
        userNotes.value = res.data.notes;
    })
    .catch(err => {
    console.error('Error fetching notes:', err);
    });
}

function logOut(){
    userService.authClient().post("/logout")
    router.push(`/`);
}

onMounted(() => {
    getMyNotes();
});

</script>
