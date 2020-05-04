window.addEventListener("load", () => {
    let blockTransactions = document.querySelector("div#block-transactions");
    let block = document.querySelector("div#block-preview");
    let blockInfo = {
        hash: block.querySelector("span#hash"),
        prevHash: block.querySelector("span#prevHash"),
        nextHash: block.querySelector("span#nextHash"),
        timestamp: block.querySelector("span#timestamp"),
        height: block.querySelector("span#height"),
        txCount: block.querySelector("span#tx-count"),
        difficulty: block.querySelector("span#difficulty"),
        merkleRoot: block.querySelector("span#merkle-root"),
        version: block.querySelector("span#version"),
        bits: block.querySelector("span#bits"),
        nonce: block.querySelector("span#nonce"),
        confirmations: block.querySelector("span#confirmations"),
    };
    let maxHeight;
    let currentHeight;

    // Populate blockchain preview
    function initPreviewWithLatest() {
        fetch(`http://localhost:8080/getLatestBlock`)
            .then(response => response.json())
            .then(data => {
                maxHeight = data.Height;
                currentHeight = data.Height;
                updateBlockPreview(data.Hash);
            });
    }

    initPreviewWithLatest();

    document.querySelector("button#forward").addEventListener("click", moveLeftOfBlock);
    document.querySelector("button#back").addEventListener("click", moveRightOfBlock);

    function moveLeftOfBlock() {
        if (currentHeight + 1 < maxHeight) {
            currentHeight++;
            fetch(`http://localhost:8080/getBlockHashFromHeight/${currentHeight + 1}`)
                .then(response => response.json())
                .then(data => updateBlockPreview(data.Hash));
        }
    }

    function moveRightOfBlock() {
        if (currentHeight - 1 > 0) {
            currentHeight--;
            fetch(`http://localhost:8080/getBlockHashFromHeight/${currentHeight - 1}`)
                .then(response => response.json())
                .then(data => updateBlockPreview(data.Hash));
        }
    }

    document.querySelector("button#search-submit").addEventListener("click", onSearchSubmit);

    function onSearchSubmit() {
        let params = document.querySelector("input#block-search").value;

        if (!isNaN(params)) {
            // Search params are block height
            fetch(`http://localhost:8080/getBlockHashFromHeight/${parseInt(params, 10)}`)
                .then(response => response.json())
                .then(data => {
                    updateBlockPreview(data.Hash);
                });
        } else {
            // Search params are block hash
            fetch(`http://localhost:8080/getBlockHeightFromHash/${parseInt(params, 10)}`)
                .then(response => response.json())
                .then(data => {
                    updateBlockPreview(data.Height);
                });
        }
    }

    function updateBlockPreview(blockHash) {
        fetch(`http://localhost:8080/getBlockHeaderFromHash/${blockHash}`)
            .then(response => response.json())
            .then(data => {
                blockInfo.hash.innerHTML = data.Hash;
                blockInfo.height.innerHTML = data.Height;
                blockInfo.confirmations.innerHTML = data.Confirmations;
                blockInfo.difficulty.innerHTML = data.Difficulty;
                blockInfo.nextHash.innerHTML = data.NextBlock;
                blockInfo.prevHash.innerHTML = data.PrevBlock;
                blockInfo.timestamp.innerHTML = data.Timestamp;
                blockInfo.merkleRoot.innerHTML = data.MerkleRoot;
                blockInfo.version.innerHTML = data.Version;
                blockInfo.bits.innerHTML = data.Bits;
                blockInfo.nonce.innerHTML = data.Nonce;
            });

        fetch(`http://localhost:8080/getBlockTransactions/${blockHash}`)
            .then(response => response.json())
            .then(data => {
                blockInfo.txCount.innerHTML = data.amount;

                let container = blockTransactions.querySelector("div#transactions-container");
                container.innerHTML = "";
                let unorderedList = document.createElement("ul");
                container.appendChild(unorderedList);

                data.transactions.forEach((tx) => {
                    let listElement = document.createElement("li");
                    listElement.innerHTML = tx;

                    unorderedList.appendChild(listElement);
                });
            });
    }

});