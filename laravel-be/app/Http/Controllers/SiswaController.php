<?php

// app/Http/Controllers/SiswaController.php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Models\Siswa;

class SiswaController extends Controller
{
    public function index()
    {
        $siswas = Siswa::all();
        return view('siswas.index', compact('siswas'));
    }

    public function create()
    {
        return view('siswas.create');
    }

    public function store(Request $request)
    {
        $request->validate([
            'nama' => 'required',
            'alamat' => 'required',
            'email' => 'required|email|unique:siswas',
        ]);

        Siswa::create($request->all());

        return redirect()->route('siswas.index')->with('success', 'Siswa created successfully.');
    }

    public function show(Siswa $siswa)
    {
        return view('siswas.show', compact('siswa'));
    }

    public function edit(Siswa $siswa)
    {
        return view('siswas.edit', compact('siswa'));
    }

    public function update(Request $request, Siswa $siswa)
    {
        $request->validate([
            'nama' => 'required',
            'alamat' => 'required',
            'email' => 'required|email|unique:siswas,email,' . $siswa->id,
        ]);

        $siswa->update($request->all());

        return redirect()->route('siswas.index')->with('success', 'Siswa updated successfully.');
    }

    public function destroy(Siswa $siswa)
    {
        $siswa->delete();
        return redirect()->route('siswas.index')->with('success', 'Siswa deleted successfully.');
    }
}
