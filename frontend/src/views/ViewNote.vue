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

<!-- <script>
import axios from 'axios';
export default{
    name:'ViewNote',
    data(){
        return {
            msg: ""
        }
    },
    methods:{
        getNoteText(){
            const path ='http://localhost:8080/note/:url';
            axios.get(path)
            .then((res)=>{
                console.log(res.data);
                this.msg= res.data
            })
            .catch ((err)=>{
                console.err(err);
            })
        },
    },
    created(){
        this.getNoteText();
    }
}
</script> -->

<script setup>
import { ref,onMounted } from 'vue';
const msg=ref("")

async function getNoteText(){
    await baseClient().get("/note/:url")//get it from msh3arfa and remove then catch
    .then((res)=>{
        console.log(res.data);
        msg.value= res.data
            })
    .catch ((err)=>{
        console.error(err);
    })
}

onMounted(() => {
  getNoteText();
});

</script>