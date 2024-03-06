<!-- resources/views/siswas/create.blade.php -->

<h1>Create Siswa</h1>

@if ($errors->any())
    <div>
        <ul>
            @foreach ($errors->all() as $error)
                <li>{{ $error }}</li>
            @endforeach
        </ul>
    </div>
@endif

<form action="{{ route('siswas.store') }}" method="post">
    @csrf
    <label for="nama">Nama:</label>
    <input type="text" name="nama" id="nama" value="{{ old('nama') }}" required>

    <label for="alamat">Alamat:</label>
    <textarea name="alamat" id="alamat" required>{{ old('alamat') }}</textarea>

    <label for="email">Email:</label>
    <input type="email" name="email" id="email" value="{{ old('email') }}" required>

    <button type="submit">Create Siswa</button>
</form>
