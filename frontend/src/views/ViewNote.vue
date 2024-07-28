<template>
    <div>
        {{ msg }}
    </div>
</template>

<script>
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
</script>