<template>
    <v-app>
        <v-container>
            <v-row class="justify-end">
                <v-btn color="primary" @click="logout">
                    <img src="/root/Secret-noteAPI-SPA-NabilaSherif/frontend/src/assets/logout.png" width=20px>
                    Log Out
                </v-btn>
            </v-row>
            <v-row>
                <v-card>
                    <h2 class="blue--text text--darken-2 display-2 font-weight-bold">Post a new card</h2>
                    <v-form @submit.prevent="onSubmitPost">
                    <v-text-field v-model="createNewNote.notetext" label="Note Text" />
                    <v-text-field v-model="createNewNote.expirationdate" label="Expiration Date" type="date" />
                    <v-text-field v-model="createNewNote.maxviews" label="Max Views" type="number" />
                    <v-card-actions>
                        <v-spacer></v-spacer>
                        <v-btn type="submit" color="primary">Post</v-btn>
                    </v-card-actions>
                    </v-form>
                </v-card>
                </v-row>
            <v-row>
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
            </v-row>
        </v-container>
    </v-app>
  </template>
  

<script>
import axios from 'axios';
import { defineComponent } from 'vue';

export default defineComponent({
    name: 'Profile',
    data() {
        return {
            createNewNote: {
                notetext: "",
                expirationdate: "",
                maxviews: "",
            },
            userNotes: {
                currentusernotes: []
            }
        };
    },
    methods: {
        initForm() {
            this.createNewNote.notetext = "";
            this.createNewNote.expirationdate = "";
            this.createNewNote.maxviews = "";
        },
        postNote() {
            const path = 'http://localhost:8080/note';
            const noteInfo = {
                notetext: this.createNewNote.notetext,
                expirationdate: this.createNewNote.expirationdate,
                maxviews: this.createNewNote.maxviews
            };

            axios.post(path, noteInfo)
                .then(response => {
                    console.log(response.data);
                    this.initForm();
                    //a success message here
                })
                .catch(err => {
                    console.error('Error creating note:', err);
                });
        },
        getMyNotes() {
            const path = 'http://localhost:8080/notes'; 
            //elmafroud yegeb mn el credientials el username
            axios.get(path)
            .then(res => {
             this.notes = res.data.notes;
            })
            .catch(err => {
            console.error('Error fetching notes:', err);
             });
        },
        logOut(){
            const path ='http://localhost:8080/logout'; 
            axios.post(path)
        },
        onSubmitPost(e){
            e.preventDefault()
            this.postNote()
            this.initForm()
        }
    },
    created() {
    this.getMyNotes();
  }
});
</script>
