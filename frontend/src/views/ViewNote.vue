<template>
    <v-app>
        <v-container>
            <v-row class="text-center">
                <v-col cols="12" class="d-flex justify-center">
                    <p>
                        {{ msg }}
                    </p>
                </v-col>
            </v-row>
        </v-container>
    </v-app>
</template>

<script setup>
import { ref,onMounted } from 'vue';
import userservice from '../services/userservice';
import { useRoute } from 'vue-router';

const router=useRoute();

const msg=ref("");

async function getNoteText(){
    const noteUrl = router.params.note_url; // Extract parameter from route
    await userservice.baseClient().get(`/note/${noteUrl}`)
    //await userservice.baseClient().get(`/note/${note_url}`)
    .then((res)=>{
        console.log(res.data);
        msg.value= res.data.note_text;
            })
    .catch ((err)=>{
        console.error(err);
    })
}

onMounted(() => {
  getNoteText();
});

</script>