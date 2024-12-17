<template>
    <v-container>
        <v-row class="justify-end">
            <v-btn color="info" @click="logOut">
                <img src="../assets/logout.png" width=20px>
                Log Out
            </v-btn>
        </v-row>
        <v-row style="padding-left: 5cm;">
            <v-col>
                <v-card max-width="800px" >
                    <v-card-title>POST NEW NOTE</v-card-title>
                    <v-form @submit.prevent="postNote" >
                        <v-card-subtitle>
                            <v-text-field v-model="note.note_text" label="Note Text" />
                            <v-text-field v-model="note.expiration_date" label="Expiration Date" type="date" />
                            <v-text-field v-model="note.max_viewers" label="Max Views" type="number" />
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
            <v-col>
                <h2 class="blue--text text--darken-2 display-2 font-weight-bold ">MY NOTES 📝</h2>
            </v-col>
        </v-row>
        <v-row style="padding-left: 5cm;">
            <v-col>
                <v-card v-for="note in notes" :key="note.note_url" class="mb-4" max-width="800px">
                    <v-card-title>{{ note.note_text }}</v-card-title>
                    <v-card-subtitle>
                        <strong>Current Views:</strong> {{ note.current_viewers }}<br>
                        <strong>Max Views:</strong> {{ note.max_viewers }}<br>
                        <strong>Expiration Date:</strong> {{ note.expiration_date }}<br>
                        <strong>URL:</strong>
                        <a :href="`http://localhost:5173/note/${note.note_url}`" target="_blank">
                            http://localhost:5173/note/{{ note.note_url }}
                        </a>
                    </v-card-subtitle>
                </v-card>
            </v-col>
        </v-row>
    </v-container>
</template>

<script setup>
import {ref, onMounted} from 'vue';
import userservice from '../services/userservice.ts';
import { useRouter } from 'vue-router';


const router =useRouter();

const showCreate=ref(true);

const note=ref( {
    note_text: "",
    expiration_date:"",
    max_viewers:"",
});

const notes=ref([])

function initForm() {
    note.value.note_text = "";
    note.value.expiration_date = "";
    note.value.max_viewers = "";
}

function postNote(){
    console.log("heree",note.value);
    userservice.postNote(note.value)
    .then(response=>{
        getMyNotes();
    })
    .catch(err=>{
        console.error('Error creating note:', err);
    });
    initForm();
}

async function  getMyNotes() {
    await userservice.authClient().get("/notes")
    .then(res => {
        console.log("Full response:", res);  // Log the full response
        console.log("Response data:", res.data);  // Log the data part of the response
        console.log("Fetched notes:", res.data.notes);  // Log the specific part of the response
        notes.value = res.data.notes;
    })
    .catch(err => {
    console.error('Error fetching notes:', err);
    });
}

function logOut(){
    userservice.logout();
    router.push('/');
}

onMounted(() => {
    getMyNotes();
});

</script>
