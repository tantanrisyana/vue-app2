<!-- resources/views/siswas/edit.blade.php -->

<h1>Edit Siswa</h1>

@if ($errors->any())
    <div>
        <ul>
            @foreach ($errors->all() as $error)
                <li>{{ $error }}</li>
            @endforeach
        </ul>
    </div>
@endif

<form action="{{ route('siswas.update', $siswa->id) }}" method="post">
    @csrf
    @method('PUT')
    <label for="nama">Nama:</label>
    <input type="text" name="nama" id="nama" value="{{ old('nama', $siswa->nama) }}" required>

    <label for="alamat">Alamat:</label>
    <textarea name="alamat" id="alamat" required>{{ old('alamat', $siswa->alamat) }}</textarea>

    <label for="email">Email:</label>
    <input type="email" name="email" id="email" value="{{ old('email', $siswa->email) }}" required>

    <button type="submit">Update Siswa</button>
</form>
<a href="{{ route('siswas.index') }}">Back to Index</a>
