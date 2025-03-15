async function sortArray() {
    const arrayInput = document.getElementById("arrayInput").value;
    const algorithm = document.getElementById("algorithm").value;
    const sortButton = document.querySelector("button");

    if (!arrayInput) {
        alert("Masukkan list angka terlebih dahulu!");
        return;
    }

    const array = arrayInput.split(",").map(num => parseInt(num.trim(), 10));
    const validAlgorithms = ["bogo", "selection"];
    sortButton.disabled = true;

    if (validAlgorithms.includes(algorithm)) {
        try {
            const response = await fetch(`/${algorithm}`, {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ array })
            });

            if (!response.ok) {
                throw new Error("Gagal melakukan sorting");
            }

            const data = await response.json();
            document.getElementById("iterations").textContent = "Iterations: " + data.iteration;
            document.getElementById("sortedArray").textContent = "Sorted Array: " + data.sorted.join(", ");
        } catch (error) {
            alert("Terjadi kesalahan: " + error.message);
        }
    } else {
        alert("Algoritma belum tersedia!");
    }

    sortButton.disabled = false;
}
