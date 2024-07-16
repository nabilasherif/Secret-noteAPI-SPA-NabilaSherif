<template>
    <div>
        <button type="button" class = "btn btn-primary"></button>
        {{ msg }}
    </div>
</template>

<script>
import axios from 'axios';
export default{
    name:'Note',
    data(){
        return {
            msg: "note text to be sent here"
        }
    },
    methods:{
        getNoteText(){
            //pass the path
            const path ='';
            axios.get(path)
            .then((res)=>{
                console.log(res.data);
            })
            .catch ((err)=>{
                console.err(err);
            })
        },
    },
    created
}
</script>