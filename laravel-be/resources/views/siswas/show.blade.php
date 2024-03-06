<!-- resources/views/siswas/show.blade.php -->

<h1>Show Siswa</h1>
<p>Nama: {{ $siswa->nama }}</p>
<p>Alamat: {{ $siswa->alamat }}</p>
<p>Email: {{ $siswa->email }}</p>
<a href="{{ route('siswas.index') }}">Back to Index</a>
