<script>
    let canSayHello = false

    function sayHello () {
        canSayHello = true
    }

</script>


{#if canSayHello} 
    <h1>Hello world</h1>    
{/if}

<button on:click={sayHello}>Hello</button>