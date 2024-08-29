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

let msg=ref("");


// async function getNoteText(){
//     const noteUrl = router.params.note_url;
//     await userservice.baseClient().get(`/note/${noteUrl}`)
//     .then((response)=>{
//         console.log(res.data);
//         msg.value= res.data.note_text;
//             })
//     .catch ((err)=>{
//         console.error(err);
//     })
// }

async function getNoteText() {
    const noteUrl = router.params.note_url;
    try {
        const response = await userservice.baseClient().get(`/note/${noteUrl}`);
        console.log('Response data:', response.data);
        msg.value = response.data.note_text;
    } catch (err) {
        console.error('Error fetching note:', err);
    }
}


onMounted(() => {
  getNoteText();
});

</script>