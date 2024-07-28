<template>
    <v-container>
        <v-row class="justify-end">
            <v-btn color="info" @click="logout">
                <img src="/root/Secret-noteAPI-SPA-NabilaSherif/frontend/src/assets/logout.png" width=20px>
                Log Out
            </v-btn>
        </v-row>

        <v-row  align-center justify="center">
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
        <!-- <v-row>
            <h2 class="blue--text text--darken-2 display-2 font-weight-bold">My Notes</h2>
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
        </v-row> -->
    </v-container>
</template>

<script setup>
import {ref} from 'vue';
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
    this.createNewNote.notetext = "";
    this.createNewNote.expirationdate = "";
    this.createNewNote.maxviews = "";
}

function onSubmitPost(){
    postNote()
    initForm()
}

async function postNote() {
    const noteInfo = {
        notetext: this.createNewNote.notetext,
        expirationdate: this.createNewNote.expirationdate,
        maxviews: this.createNewNote.maxviews
    };
    await baseClient().post("/note", noteInfo)
    .then(response => {
            console.log(response.data);
            toast.success('Note created successfully');
            this.initForm();
    })
    .catch(err => {
        toast.success('Error creating note');
        console.error('Error creating note:', err);
    });
}

// function  getMyNotes() {
//     //elmafroud yegeb mn el credientials el username
//     baseClient().get("/notes")
//     .then(res => {
//         this.notes = res.data.notes;
//     })
//     .catch(err => {
//     console.error('Error fetching notes:', err);
//         });
// }

function logOut(){
    baseClient().post("/logout")
}

// onMounted(() => {
//     getMyNotes();
// });

</script>
